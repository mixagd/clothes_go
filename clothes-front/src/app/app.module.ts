import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';


import {AppComponent} from './app.component';
import {HomePageComponent} from './home-page/home-page.component';
import {ClothesService} from './clothes.service';
import {HttpClientModule} from '@angular/common/http';
import {ClothItemComponent} from './cloth-item/cloth-item.component';
import {appRoutes} from './/app-routing.module';
import {ClothesComponent} from './clothes/clothes.component';
import {RouterModule} from '@angular/router';
import {ClothFormComponent} from './cloth-form/cloth-form.component';
import {FormsModule} from '@angular/forms';


@NgModule({
  declarations: [
    AppComponent,
    HomePageComponent,
    ClothItemComponent,
    ClothesComponent,
    ClothFormComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    RouterModule.forRoot(appRoutes),
    FormsModule,
    // AppRouting
  ],
  providers: [ClothesService],
  bootstrap: [AppComponent]
})
export class AppModule { }
