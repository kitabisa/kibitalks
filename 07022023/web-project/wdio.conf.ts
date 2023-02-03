const { TimelineService } = require('wdio-timeline-reporter/timeline-service');

export const config: WebdriverIO.Config = {
  // Runner Configuration
  runner: 'local',

  autoCompileOpts: {
    tsNodeOpts: {
      project: './tsconfig.json'
    }
  },
  
  specs: ['./test/specs/*.ts'],
  exclude: [],

  // Capabilities
  maxInstances: 10,
  capabilities: [{
    maxInstances: 1,
    browserName: 'chrome',
    acceptInsecureCerts: true,
    "goog:chromeOptions": {
      args: [
        "--headless",
        "--incognito",
      ]
      }
  }],

  // Test Configurations
  logLevel: 'error',
  bail: 0,
  baseUrl: "https://www.facebook.com/signup",
  waitforTimeout: 30000,
  connectionRetryTimeout: 40000,
  connectionRetryCount: 2,
  services: [[TimelineService], 'chromedriver'],
  framework: 'mocha',
  reporters: [ "spec", ['timeline', 
  { 
    outputDir: './reporter',
    embedImages: true,
    images: {
      quality: 80,
      resize: false,
      reductionRatio: 1
    },
    screenshotStrategy: 'on:error'
  }] 
],
  
  mochaOpts: {
    ui: 'bdd',
    timeout: 40000
  },
}
