import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ProcedureOrderService } from '../../../services/ProcedureOrder.service';
import { ProcedureOrder } from '../../../models/ProcedureOrder';
import { SubBaseComponent } from '../../ProcedureOrder/sub.base.component';

@Component({
    selector: 'app-create-procedureOrder',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateProcedureOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Add ProcedureOrder';

    procedureOrderForm: FormGroup;
    procedureOrder: ProcedureOrder;

    constructor( http: HttpClient,
        private procedureOrderService: ProcedureOrderService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.procedureOrderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  procedureCode: ['', Validators.required],
      consentObtained: ['', Validators.required],
      Order: ['', ],
      Facility: ['', ],
      Procedure: ['', ],
      AnesthesiaType: ['', ]
        });
    }

    
    addProcedureOrder(procedureCode, consentObtained, Order, Facility, Procedure, AnesthesiaType): void {
        this.procedureOrderService
        .addProcedureOrder(procedureCode, consentObtained, Order, Facility, Procedure, AnesthesiaType)
            .subscribe(() => {
                this.router.navigate(['/indexProcedureOrder']);
            });
    }

    ngOnInit(): void {
    }
}