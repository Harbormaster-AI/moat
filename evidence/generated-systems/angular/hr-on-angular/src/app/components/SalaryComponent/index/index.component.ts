
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { SalaryComponentService } from '../../../services/SalaryComponent.service';
import { SalaryComponent } from '../../../models/SalaryComponent';

@Component({
    selector: 'app-index-salaryComponent',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexSalaryComponentComponent implements OnInit {

    salaryComponents: SalaryComponent[] = [];

    constructor(
        private router: Router,
        private service: SalaryComponentService
) {}

    ngOnInit(): void {
        this.getSalaryComponents();
}

    getSalaryComponents(): void {
        this.service.getSalaryComponents().subscribe((res) => {
        this.salaryComponents = res;
    });
}

    deleteSalaryComponent(id: any): void {
        this.service.deleteSalaryComponent(id)
            .subscribe(() => {
                this.getSalaryComponents();
            });
    }
}