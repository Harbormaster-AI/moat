
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ProcedureOrderService } from '../../../services/ProcedureOrder.service';
import { ProcedureOrder } from '../../../models/ProcedureOrder';

@Component({
    selector: 'app-index-procedureOrder',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexProcedureOrderComponent implements OnInit {

    procedureOrders: ProcedureOrder[] = [];

    constructor(
        private router: Router,
        private service: ProcedureOrderService
) {}

    ngOnInit(): void {
        this.getProcedureOrders();
}

    getProcedureOrders(): void {
        this.service.getProcedureOrders().subscribe((res) => {
        this.procedureOrders = res;
    });
}

    deleteProcedureOrder(id: any): void {
        this.service.deleteProcedureOrder(id)
            .subscribe(() => {
                this.getProcedureOrders();
            });
    }
}