
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { EnterpriseService } from '../../../services/Enterprise.service';
import { Enterprise } from '../../../models/Enterprise';

@Component({
    selector: 'app-index-enterprise',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexEnterpriseComponent implements OnInit {

    enterprises: Enterprise[] = [];

    constructor(
        private router: Router,
        private service: EnterpriseService
) {}

    ngOnInit(): void {
        this.getEnterprises();
}

    getEnterprises(): void {
        this.service.getEnterprises().subscribe((res) => {
        this.enterprises = res;
    });
}

    deleteEnterprise(id: any): void {
        this.service.deleteEnterprise(id)
            .subscribe(() => {
                this.getEnterprises();
            });
    }
}