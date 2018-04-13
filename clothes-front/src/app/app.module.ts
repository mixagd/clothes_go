import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';


import { AppComponent } from './app.component';
import { HomePageComponent } from './home-page/home-page.component';
import {ClothesService} from './clothes.service';
import {HttpClientModule} from '@angular/common/http';
import { ClothItemComponent } from './cloth-item/cloth-item.component';
import {AppRouting} from './/app-routing.module';
import { ClothesComponent } from './clothes/clothes.component';


@NgModule({
  declarations: [
    AppComponent,
    HomePageComponent,
    ClothItemComponent,
    ClothesComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRouting
  ],
  providers: [ClothesService],
  bootstrap: [AppComponent]
})
export class AppModule { }
