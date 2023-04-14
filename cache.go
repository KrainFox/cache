package cache

type cache map[string]interface{}

func New()cache{
	return cache{}
}

func (c cache) Set(key string,value interface{}){
	c[key] = value
}

func (c cache) Get(key string)interface{}{
	return c[key]
}

func (c cache) Delete(key string){
	delete(c,key)
}
