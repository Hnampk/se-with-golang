package solid

/*
Make the sword blunt after a number of attacks
Cause the player to drop the sword if the monster uses some special armor
*/
// Sharpen increases the damage dealt by this sword using a whetstone.
func (Sword) Sharpen() {
	//...
}

// MakeBlunt decreases the damage dealt by this sword due to constant use.
func (Sword) MakeBlunt() {
	//...
}

// Drop places the sword on the ground allowing others to pick it up.
func (Sword) Drop() {
	//...
}

// // Attack deals damage to a monster using a sword.
// func Attack(m *Monster, s *Sword) {
// 	//...
// }

/* FIXED:
- Attack must be able to work with other types of weapons, for example, projectiles or magic spells.
- the player should be able to use weapons to deal damage to non-monster entities, for example, to break down
a bolted door or to cut the rope suspending the chandelier from the ceiling.
- reuse the same implementation when monsters attack the player
*/

// DamageReceiver is implemented by objects that can receive weapon damage.
type DamageReceiver interface {
	ApplyDamage(int)
}

// Damager is implemented by objects that can be used as weapons.
type Damager interface {
	Damage(int)
}

// Attack deals weapon damage to target.
func Attack(target DamageReceiver, weapon Damager) {
	//...
}
