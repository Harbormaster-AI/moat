
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AerospaceManufacturerService } from '../../../services/AerospaceManufacturer.service';
import { AerospaceManufacturer } from '../../../models/AerospaceManufacturer';

@Component({
    selector: 'app-index-aerospaceManufacturer',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAerospaceManufacturerComponent implements OnInit {

    aerospaceManufacturers: AerospaceManufacturer[] = [];

    constructor(
        private router: Router,
        private service: AerospaceManufacturerService
) {}

    ngOnInit(): void {
        this.getAerospaceManufacturers();
}

    getAerospaceManufacturers(): void {
        this.service.getAerospaceManufacturers().subscribe((res) => {
        this.aerospaceManufacturers = res;
    });
}

    deleteAerospaceManufacturer(id: any): void {
        this.service.deleteAerospaceManufacturer(id)
            .subscribe(() => {
                this.getAerospaceManufacturers();
            });
    }
}