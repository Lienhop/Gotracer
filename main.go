package main

func main() {
	world := hittableList{}

	world.add(Sphere{center: Vec3{0, -100.5, -1}, radius: 100})
	world.add(Sphere{center: Vec3{0, 0, -1}, radius: 0.5})
	//world.add(Sphere{center: Vec3{-2, 0, -1}, radius: 0.5})

	cam := Camera{}
	cam.render(world)
}
