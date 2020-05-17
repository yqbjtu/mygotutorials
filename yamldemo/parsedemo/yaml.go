package main
// Yaml struct of yaml
type Yaml struct {
	App struct {
		Name string `yaml:"name"`
	}
    Mysql struct {
        Host string `yaml:"host"`
		Port int32 `yaml:"port"`
		DbName string `yaml:"dbName"`
		User string `yaml:"user"`
        Password string `yaml:"password"`
    }
    Cache struct {
        Enable bool `yaml:"enable"`
        List []string `yaml:"list,flow"`
    }
}
