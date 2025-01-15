
def comparar_arquivos(arquivo1, arquivo2):
    try:
        # Abre os arquivos no modo de leitura
        with open(arquivo1, 'r', encoding='utf-8') as f1, open(arquivo2, 'r', encoding='utf-8') as f2:
            # Lê as linhas dos arquivos
            linhas1 = f1.readlines()
            linhas2 = f2.readlines()
            
            # Verifica se os arquivos têm o mesmo número de linhas
            if len(linhas1) != len(linhas2):
                print("Os arquivos têm um número diferente de linhas.")
            
            # Compara linha por linha
            for i, (linha1, linha2) in enumerate(zip(linhas1, linhas2)):
                if linha1 != linha2:
                    print(f"Diferença encontrada na linha {i + 1}:")
                    print(f"Arquivo 1: {linha1.strip()}")
                    print(f"Arquivo 2: {linha2.strip()}")
                    print("-" * 50)
                    
            # Verifica se algum arquivo tem linhas extras
            if len(linhas1) > len(linhas2):
                print(f"Arquivo 1 tem {len(linhas1) - len(linhas2)} linha(s) extra(s) no final.")
            elif len(linhas2) > len(linhas1):
                print(f"Arquivo 2 tem {len(linhas2) - len(linhas1)} linha(s) extra(s) no final.")
            
    except FileNotFoundError:
        print("Um ou ambos os arquivos não foram encontrados.")
    except Exception as e:
        print(f"Ocorreu um erro: {e}")


# Exemplo de uso
arquivo1 = "./simples/out.txt"
arquivo2 = "./simples/outbee.txt"

comparar_arquivos(arquivo1, arquivo2)
