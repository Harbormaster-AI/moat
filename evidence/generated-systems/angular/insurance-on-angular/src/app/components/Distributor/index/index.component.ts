
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DistributorService } from '../../../services/Distributor.service';
import { Distributor } from '../../../models/Distributor';

@Component({
    selector: 'app-index-distributor',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexDistributorComponent implements OnInit {

    distributors: Distributor[] = [];

    constructor(
        private router: Router,
        private service: DistributorService
) {}

    ngOnInit(): void {
        this.getDistributors();
}

    getDistributors(): void {
        this.service.getDistributors().subscribe((res) => {
        this.distributors = res;
    });
}

    deleteDistributor(id: any): void {
        this.service.deleteDistributor(id)
            .subscribe(() => {
                this.getDistributors();
            });
    }
}