import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ProcedureOrderService } from '../../../services/ProcedureOrder.service';
import { SubBaseComponent } from '../../ProcedureOrder/sub.base.component';


@Component({
    selector: 'app-edit-procedureOrder',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditProcedureOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ProcedureOrder';

    procedureOrderForm: FormGroup;
    procedureOrder: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ProcedureOrderService,
        private fb: FormBuilder
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

    
    updateProcedureOrder(procedureCode, consentObtained, Order, Facility, Procedure, AnesthesiaType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateProcedureOrder(procedureCode, consentObtained, Order, Facility, Procedure, AnesthesiaType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexProcedureOrder']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getProcedureOrder(params['id']).subscribe(res => {
                this.procedureOrder = res;
            });
        });
    }
}