package service

var keyValStore = map[string]string{}

type InMemoryService struct {
	fileService FileService
}

func NewInMemoryService(fileService FileService) InMemoryService {
	return InMemoryService{fileService}
}

func (p *InMemoryService) SetKey(key string, value string) string {
	keyValStore[key] = value
	message := key + " is " + value

	return message
}

func (p *InMemoryService) GetKey(key string) string {
	return keyValStore[key]
}

func (p *InMemoryService) GetAllKey() map[string]string {
	return keyValStore
}

func (p *InMemoryService) SetStoreToMemory() bool {
	vals := p.fileService.ReadFile();
	keyValStore = vals

	return true
}

func (p *InMemoryService) SetMemoryToStore() bool {
	_ = p.fileService.WriteFile(keyValStore);
	return true
}
