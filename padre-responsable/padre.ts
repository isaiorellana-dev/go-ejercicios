
const darPistoAHijo = (hijoName: string, cantidad: number) => {
  const index = hijos.findIndex((hijo) => hijo.name === hijoName)

  const hijosQueSePortaronBien = hijos.filter(
    (hijo) => hijo.sePortoBien === true && hijo.name != hijoName
  )
  const hijosQueSePortaronMal = hijos.filter(
    (hijo) => hijo.sePortoBien === false
  )

  const dinero = 0
  const dineroDeLosqueSePortaronBien = hijosQueSePortaronBien.reduce(
    (previus, curr) => previus + curr.mesada,
    dinero
  )

  const dineroParaDarleAlHijo = pisto - dineroDeLosqueSePortaronBien

  if (cantidad <= dineroParaDarleAlHijo) {
    hijos[index].mesada = cantidad

    const dineroARepartirALosPlebeyos =
      pisto - dineroDeLosqueSePortaronBien - cantidad

    hijos.forEach((hijo) => {
      if (!hijo.sePortoBien)
        hijo.mesada = dineroARepartirALosPlebeyos / hijosQueSePortaronMal.length
    })

    return hijos
  } else {
    return "No ajusta la plata, pida menos"
  }
}