
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CompensationPackageService } from '../../../services/CompensationPackage.service';
import { CompensationPackage } from '../../../models/CompensationPackage';

@Component({
    selector: 'app-index-compensationPackage',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCompensationPackageComponent implements OnInit {

    compensationPackages: CompensationPackage[] = [];

    constructor(
        private router: Router,
        private service: CompensationPackageService
) {}

    ngOnInit(): void {
        this.getCompensationPackages();
}

    getCompensationPackages(): void {
        this.service.getCompensationPackages().subscribe((res) => {
        this.compensationPackages = res;
    });
}

    deleteCompensationPackage(id: any): void {
        this.service.deleteCompensationPackage(id)
            .subscribe(() => {
                this.getCompensationPackages();
            });
    }
}