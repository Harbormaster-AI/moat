
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { SoftwareUpdateService } from '../../../services/SoftwareUpdate.service';
import { SoftwareUpdate } from '../../../models/SoftwareUpdate';

@Component({
    selector: 'app-index-softwareUpdate',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexSoftwareUpdateComponent implements OnInit {

    softwareUpdates: SoftwareUpdate[] = [];

    constructor(
        private router: Router,
        private service: SoftwareUpdateService
) {}

    ngOnInit(): void {
        this.getSoftwareUpdates();
}

    getSoftwareUpdates(): void {
        this.service.getSoftwareUpdates().subscribe((res) => {
        this.softwareUpdates = res;
    });
}

    deleteSoftwareUpdate(id: any): void {
        this.service.deleteSoftwareUpdate(id)
            .subscribe(() => {
                this.getSoftwareUpdates();
            });
    }
}