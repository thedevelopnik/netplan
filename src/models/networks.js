import { addVPC, createMetadata } from './helpers'

export async function initializeExistingNetworks () {
  let networks = {
    vpcs: []
  }

  try {
    const awsPrdVpc = createMetadata('aws-core', 'public', 'us-east-1', 'AWS', 'prd', 'vpc', '10.10.0.0/16')
    const awsPrdSubnets = [
      createMetadata('', 'private', 'us-east-1a', 'AWS', 'prd', 'subnet', '10.10.1.0/24'),
      createMetadata('', 'private', 'us-east-1b', 'AWS', 'prd', 'subnet', '10.10.2.0/24'),
      createMetadata('', 'private', 'us-east-1c', 'AWS', 'prd', 'subnet', '10.10.3.0/24'),
      createMetadata('', 'public', 'us-east-1a', 'AWS', 'prd', 'subnet', '10.10.11.0/24'),
      createMetadata('', 'public', 'us-east-1b', 'AWS', 'prd', 'subnet', '10.10.12.0/24'),
      createMetadata('', 'public', 'us-east-1c', 'AWS', 'prd', 'subnet', '10.10.23.0/24'),
      createMetadata('', 'database', 'us-east-1a', 'AWS', 'prd', 'subnet', '10.10.21.0/24'),
      createMetadata('', 'database', 'us-east-1b', 'AWS', 'prd', 'subnet', '10.10.22.0/24'),
      createMetadata('', 'database', 'us-east-1c', 'AWS', 'prd', 'subnet', '10.10.23.0/24')
    ]
    networks = await addVPC(networks, awsPrdVpc, awsPrdSubnets)
  } catch (err) {
    throw new Error(`could not add aws vpc: ${err}`)
  }

  try {
    const databricksVpc = createMetadata('databricks', 'public', 'us-east-1', 'AWS', 'all', 'vpc', '10.169.0.0/16')
    networks = await addVPC(networks, databricksVpc)
  } catch (err) {
    throw new Error(`could not add databricks vpc: ${err}`)
  }

  try {
    const appUsEast4 = createMetadata('app-us-east4', 'private', 'us-east4', 'GCP', 'prd', 'subnet', '10.128.16.0/20')
    networks = await addVPC(networks, appUsEast4)
  } catch (err) {
    throw new Error(`could not add app us-east4: ${err}`)
  }

  try {
    const coreBravoPods = createMetadata('core-bravo-pods', 'private', 'us-east4', 'GCP', 'prd', 'subnet', '10.16.0.0/14')
    networks = await addVPC(networks, coreBravoPods)
  } catch (err) {
    throw new Error(`could not add core bravo pods: ${err}`)
  }

  try {
    const coreBravoServices = createMetadata('core-bravo-services', 'private', 'us-east4', 'GCP', 'prd', 'subnet', '10.20.0.0/20')
    networks = await addVPC(networks, coreBravoServices)
  } catch (err) {
    throw new Error(`could not add core bravo services: ${err}`)
  }

  try {
    const aivenPrdPeer = createMetadata('aiven-peer', 'private', 'us-east4', 'Aiven', 'prd', 'vpc', '192.168.0.0/24')
    networks = await addVPC(networks, aivenPrdPeer)
  } catch (err) {
    throw new Error(`could not add aiven prd peer: ${err}`)
  }

  return networks
}
