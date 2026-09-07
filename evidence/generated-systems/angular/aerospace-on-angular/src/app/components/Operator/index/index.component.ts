
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { OperatorService } from '../../../services/Operator.service';
import { Operator } from '../../../models/Operator';

@Component({
    selector: 'app-index-operator',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexOperatorComponent implements OnInit {

    operators: Operator[] = [];

    constructor(
        private router: Router,
        private service: OperatorService
) {}

    ngOnInit(): void {
        this.getOperators();
}

    getOperators(): void {
        this.service.getOperators().subscribe((res) => {
        this.operators = res;
    });
}

    deleteOperator(id: any): void {
        this.service.deleteOperator(id)
            .subscribe(() => {
                this.getOperators();
            });
    }
}