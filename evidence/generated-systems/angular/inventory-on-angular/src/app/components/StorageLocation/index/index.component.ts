
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { StorageLocationService } from '../../../services/StorageLocation.service';
import { StorageLocation } from '../../../models/StorageLocation';

@Component({
    selector: 'app-index-storageLocation',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexStorageLocationComponent implements OnInit {

    storageLocations: StorageLocation[] = [];

    constructor(
        private router: Router,
        private service: StorageLocationService
) {}

    ngOnInit(): void {
        this.getStorageLocations();
}

    getStorageLocations(): void {
        this.service.getStorageLocations().subscribe((res) => {
        this.storageLocations = res;
    });
}

    deleteStorageLocation(id: any): void {
        this.service.deleteStorageLocation(id)
            .subscribe(() => {
                this.getStorageLocations();
            });
    }
}