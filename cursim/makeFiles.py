import os

# Cria a pasta 'ninja6' se não existir
pasta = 'ninja6'
if not os.path.exists(pasta):
    os.makedirs(pasta)

# Cria os 4 arquivos Go dentro da pasta 'ninja6'
for i in range(0, 4):
    arquivo_path = os.path.join(pasta, f'atvd{i}.go')
    with open(arquivo_path, 'w') as f:
        f.write('''package main

func main(){
    
}
''')

print("Pasta 'ninja6' e arquivos Go criados com sucesso!")
