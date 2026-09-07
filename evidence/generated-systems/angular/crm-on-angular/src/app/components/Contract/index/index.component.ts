
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ContractService } from '../../../services/Contract.service';
import { Contract } from '../../../models/Contract';

@Component({
    selector: 'app-index-contract',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexContractComponent implements OnInit {

    contracts: Contract[] = [];

    constructor(
        private router: Router,
        private service: ContractService
) {}

    ngOnInit(): void {
        this.getContracts();
}

    getContracts(): void {
        this.service.getContracts().subscribe((res) => {
        this.contracts = res;
    });
}

    deleteContract(id: any): void {
        this.service.deleteContract(id)
            .subscribe(() => {
                this.getContracts();
            });
    }
}