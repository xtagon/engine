package rules

const (
	// DeathCauseSnakeCollision is the death reason when 2 snakes collide with each other
	DeathCauseSnakeCollision = "snake-collision"
	// DeathCauseSnakeSelfCollision is the death reason when a snake collides with itself
	DeathCauseSnakeSelfCollision = "snake-self-collision"
	// DeathCauseStarvation is the death reason when a snakes health reaches zero
	DeathCauseStarvation = "starvation"
	// DeathCauseHeadToHeadCollision is when a snake dies from a head on head collision
	DeathCauseHeadToHeadCollision = "head-collision"
	// DeathCauseWallCollision is when a snake runs off the board
	DeathCauseWallCollision = "wall-collision"
)
