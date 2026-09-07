
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { OperationService } from '../../../services/Operation.service';
import { Operation } from '../../../models/Operation';

@Component({
    selector: 'app-index-operation',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexOperationComponent implements OnInit {

    operations: Operation[] = [];

    constructor(
        private router: Router,
        private service: OperationService
) {}

    ngOnInit(): void {
        this.getOperations();
}

    getOperations(): void {
        this.service.getOperations().subscribe((res) => {
        this.operations = res;
    });
}

    deleteOperation(id: any): void {
        this.service.deleteOperation(id)
            .subscribe(() => {
                this.getOperations();
            });
    }
}