import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PayrollItemService } from '../../../services/PayrollItem.service';
import { SubBaseComponent } from '../../PayrollItem/sub.base.component';


@Component({
    selector: 'app-edit-payrollItem',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPayrollItemComponent extends SubBaseComponent implements OnInit {

    title = 'Edit PayrollItem';

    payrollItemForm: FormGroup;
    payrollItem: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PayrollItemService,
        private fb: FormBuilder
) {
        super(http);
        this.payrollItemForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  amount: ['', Validators.required],
      taxable: ['', Validators.required],
      PayrollRun: ['', ],
      Employee: ['', ],
      ItemType: ['', ]
        });
    }

    
    updatePayrollItem(amount, taxable, PayrollRun, Employee, ItemType): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePayrollItem(amount, taxable, PayrollRun, Employee, ItemType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPayrollItem']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPayrollItem(params['id']).subscribe(res => {
                this.payrollItem = res;
            });
        });
    }
}