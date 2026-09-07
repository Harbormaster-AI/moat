
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PlantService } from '../../../services/Plant.service';
import { Plant } from '../../../models/Plant';

@Component({
    selector: 'app-index-plant',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPlantComponent implements OnInit {

    plants: Plant[] = [];

    constructor(
        private router: Router,
        private service: PlantService
) {}

    ngOnInit(): void {
        this.getPlants();
}

    getPlants(): void {
        this.service.getPlants().subscribe((res) => {
        this.plants = res;
    });
}

    deletePlant(id: any): void {
        this.service.deletePlant(id)
            .subscribe(() => {
                this.getPlants();
            });
    }
}