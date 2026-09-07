
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { EmploymentContractService } from '../../../services/EmploymentContract.service';
import { EmploymentContract } from '../../../models/EmploymentContract';

@Component({
    selector: 'app-index-employmentContract',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexEmploymentContractComponent implements OnInit {

    employmentContracts: EmploymentContract[] = [];

    constructor(
        private router: Router,
        private service: EmploymentContractService
) {}

    ngOnInit(): void {
        this.getEmploymentContracts();
}

    getEmploymentContracts(): void {
        this.service.getEmploymentContracts().subscribe((res) => {
        this.employmentContracts = res;
    });
}

    deleteEmploymentContract(id: any): void {
        this.service.deleteEmploymentContract(id)
            .subscribe(() => {
                this.getEmploymentContracts();
            });
    }
}