package main

func main() {
	world := hittableList{}

	world.add(Sphere{center: Vec3{0, -100.5, -1}, radius: 100, material: Lambertian{albedo: Color{r: 0.8, g: 0.8, b: 0.0}}})
	world.add(Sphere{center: Vec3{0, 0, -1}, radius: 0.5, material: Metal{albedo: Color{r: 0.7, g: 0.7, b: 0.7}, fuzz: 0.3}})
	world.add(Sphere{center: Vec3{-1.2, 0, -1}, radius: 0.5, material: Lambertian{albedo: Color{r: 0.1, g: 0.1, b: 0.7}}})
	world.add(Sphere{center: Vec3{1.2, 0, -1}, radius: 0.5, material: Dielectric{refractionIndex: 1.00 / 1.33}})

	cam := Camera{}
	cam.render(world)
}
