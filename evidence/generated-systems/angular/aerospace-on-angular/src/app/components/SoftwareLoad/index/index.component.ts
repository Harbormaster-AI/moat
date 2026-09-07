
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { SoftwareLoadService } from '../../../services/SoftwareLoad.service';
import { SoftwareLoad } from '../../../models/SoftwareLoad';

@Component({
    selector: 'app-index-softwareLoad',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexSoftwareLoadComponent implements OnInit {

    softwareLoads: SoftwareLoad[] = [];

    constructor(
        private router: Router,
        private service: SoftwareLoadService
) {}

    ngOnInit(): void {
        this.getSoftwareLoads();
}

    getSoftwareLoads(): void {
        this.service.getSoftwareLoads().subscribe((res) => {
        this.softwareLoads = res;
    });
}

    deleteSoftwareLoad(id: any): void {
        this.service.deleteSoftwareLoad(id)
            .subscribe(() => {
                this.getSoftwareLoads();
            });
    }
}