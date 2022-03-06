import { NgModule } from '@angular/core';
import {PreloadAllModules, RouteReuseStrategy, RouterModule, Routes} from '@angular/router';

const harborRoutes: Routes = [
  {
    path: 'devcenter-api-2.0',
    loadChildren: () => import('./dev-center/dev-center.module').then(m => m.DeveloperCenterModule)
  },
  {
    path: 'license',
    loadChildren: () => import('./license/license.module').then(m => m.LicenseModule)
  },
  {
    path: 'harbor',
    loadChildren: () => import('./base/base.module').then(m => m.BaseModule)
  }
]

@NgModule({
  providers: [

  ],
  imports: [
    RouterModule.forRoot(harborRoutes, {
      onSameUrlNavigation: 'reload',
      preloadingStrategy: PreloadAllModules,
      relativeLinkResolution: 'legacy'
    })
  ],
  exports: [RouterModule]
})
export class HarborRoutingModule {}