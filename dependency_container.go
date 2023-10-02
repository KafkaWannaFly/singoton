package singoton

type _DependencyContainer struct {
	objectContainerMap    map[Metadata]any
	interfaceImplementMap map[Metadata]Metadata
	factoryContainerMap   map[Metadata]any
}

func (container *_DependencyContainer) addObject(metadata Metadata, object any) {
	container.objectContainerMap[metadata] = object
}

func (container *_DependencyContainer) addInterfaceImplement(interfaceMetadata Metadata, implementMetadata Metadata) {
	container.interfaceImplementMap[interfaceMetadata] = implementMetadata
}

func (container *_DependencyContainer) addFactory(metadata Metadata, factory any) {
	container.factoryContainerMap[metadata] = factory
}

func (container *_DependencyContainer) getObject(metadata Metadata) (any, bool) {
	object, ok := container.objectContainerMap[metadata]
	return object, ok
}

func (container *_DependencyContainer) getInterfaceImplement(interfaceMetadata Metadata) (Metadata, bool) {
	implementMetadata, ok := container.interfaceImplementMap[interfaceMetadata]
	return implementMetadata, ok
}

func (container *_DependencyContainer) getFactory(metadata Metadata) (any, bool) {
	factory, ok := container.factoryContainerMap[metadata]
	return factory, ok
}

func (container *_DependencyContainer) removeObject(metadata Metadata) {
	delete(container.objectContainerMap, metadata)
}

func (container *_DependencyContainer) removeInterfaceImplement(interfaceMetadata Metadata) {
	delete(container.interfaceImplementMap, interfaceMetadata)
}

func (container *_DependencyContainer) removeFactory(metadata Metadata) {
	delete(container.factoryContainerMap, metadata)
}

func (container *_DependencyContainer) getObjectMap() *map[Metadata]any {
	return &container.objectContainerMap
}

func (container *_DependencyContainer) getInterfaceImplementMap() *map[Metadata]Metadata {
	return &container.interfaceImplementMap
}

func (container *_DependencyContainer) getFactoryMap() *map[Metadata]any {
	return &container.factoryContainerMap
}
