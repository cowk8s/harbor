import { NgModule } from '@angular/core';
import { PreloadAllModules, RouterModule, Routes } from '@angular/router';

const harborRoutes: Routes = [
  {path: '', redirectTo: 'harbor', pathMatch: 'full' },
  {
    path: 'devcenter-api-2.0',
    loadChildren: () => import('./dev-center/dev-center.module').then(m => m.DeveloperCenterModule)
  },
  {
    path: '**',
    loadChildren: () => import('./not-found/not-found.module').then(m => m.NotFoundModule)
  }
];

@NgModule({
  providers: [

  ],
  imports: [
    RouterModule.forRoot(harborRoutes, {
      onSameUrlNavigation: 'reload',
      preloadingStrategy: PreloadAllModules
    })
  ],
  exports: [RouterModule]
})
export class HarborRoutingModule {}