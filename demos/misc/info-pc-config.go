//go:build ignore

package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/fatih/color"
	"github.com/jaypipes/ghw"
)

func main() {
	// Obter informações de CPU
	cpuInfo, err := ghw.CPU()
	if err != nil {
		log.Printf("Erro ao obter informações da CPU: %v", err)
	}

	// Obter informações de memória
	memInfo, err := ghw.Memory()
	if err != nil {
		log.Printf("Erro ao obter informações da memória: %v", err)
	}

	// Obter informações de dispositivos de armazenamento
	// blockInfo, err := ghw.Block()
	// if err != nil {
	// 	log.Printf("Erro ao obter informações dos discos: %v", err)
	// }

	// Configurando cores com fatih/color
	title := color.New(color.FgHiGreen, color.Bold).SprintFunc()
	label := color.New(color.FgCyan).SprintFunc()
	infoColor := color.New(color.FgWhite).SprintFunc()

	// Cabeçalho
	fmt.Println(title("⚙️  Informações do Sistema"))
	fmt.Println("------------------------------------")

	// Seção CPU
	fmt.Println(title("🖥️  CPU:"))
	if cpuInfo != nil && len(cpuInfo.Processors) > 0 {
		proc := cpuInfo.Processors[0]
		vendor := "Desconhecido"
		if proc.Vendor != "" {
			vendor = proc.Vendor
		}
		fmt.Printf("  %s %s %s\n", label("Modelo:"), infoColor(vendor), infoColor(proc.Model))
		fmt.Printf("  %s %d núcleos físicos, %d núcleos lógicos\n", label("Cores:"), proc.NumCores, proc.NumThreads)
		fmt.Printf("  %s %s\n", label("Frequência:"), infoColor("Não disponível"))
	} else {
		fmt.Println("  Informação da CPU não disponível.")
	}

	// Seção Memória
	fmt.Println(title("\n💾 Memória:"))
	if memInfo != nil {
		totalGB := float64(memInfo.TotalUsableBytes) / 1e9
		fmt.Printf("  %s %.2f GB\n", label("Total:"), totalGB)
		// A frequência da memória pode não estar disponível via ghw
		fmt.Printf("  %s %s\n", label("Frequência:"), infoColor("Não disponível"))
	} else {
		fmt.Println("  Informação da memória não disponível.")
	}

	// Seção SSD
	// fmt.Println(title("\n📀 Armazenamento (SSD):"))
	// if blockInfo != nil {
	// 	foundSSD := false
	// 	for _, disk := range blockInfo.Disks {
	// 		// No WSL e outros ambientes, podemos precisar verificar diferentes tipos
	// 		if disk.DriveType == ghw.DRIVE_TYPE_SSD || disk.DriveType == ghw.DRIVE_TYPE_NVME {
	// 			foundSSD = true
	// 			model := disk.Model
	// 			if model == "" {
	// 				model = "Dispositivo desconhecido"
	// 			}
	// 			fmt.Printf("  %s %s\n", label("Modelo:"), infoColor(model))
	// 			capGB := float64(disk.SizeBytes) / 1e9
	// 			fmt.Printf("  %s %.2f GB\n", label("Capacidade:"), capGB)
	// 			fmt.Printf("  %s %s\n", label("Velocidade:"), infoColor("Não disponível"))
	// 			fmt.Println("")
	// 		}
	// 	}
	// 	if !foundSSD {
	// 		fmt.Println("  Nenhum SSD encontrado.")
	// 	}
	// } else {
	// 	fmt.Println("  Informação do armazenamento não disponível.")
	// }

	// Seção GPU
	fmt.Println(title("🎮 GPU:"))
	pciInfo, err := ghw.PCI()
	if err != nil {
		log.Printf("Erro ao obter informações PCI: %v", err)
	} else {
		foundGPU := false
		for _, device := range pciInfo.Devices {
			// Classe 0x03 é para controladores de vídeo/gráficos
			classID, err := strconv.ParseInt(device.Class.ID, 16, 64)
			if err == nil && classID == 0x03 {
				foundGPU = true
				fmt.Printf("  %s %s\n", label("Modelo:"), infoColor(device.Product.Name))
			}
		}
		if !foundGPU {
			fmt.Println("  Informação da GPU não disponível.")
		}
	}

	fmt.Println("------------------------------------")
}
