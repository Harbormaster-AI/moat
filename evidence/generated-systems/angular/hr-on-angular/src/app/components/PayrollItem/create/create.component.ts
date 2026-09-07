import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PayrollItemService } from '../../../services/PayrollItem.service';
import { PayrollItem } from '../../../models/PayrollItem';
import { SubBaseComponent } from '../../PayrollItem/sub.base.component';

@Component({
    selector: 'app-create-payrollItem',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePayrollItemComponent extends SubBaseComponent implements OnInit {

    title = 'Add PayrollItem';

    payrollItemForm: FormGroup;
    payrollItem: PayrollItem;

    constructor( http: HttpClient,
        private payrollItemService: PayrollItemService,
        private fb: FormBuilder,
        private router: Router
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

    
    addPayrollItem(amount, taxable, PayrollRun, Employee, ItemType): void {
        this.payrollItemService
        .addPayrollItem(amount, taxable, PayrollRun, Employee, ItemType)
            .subscribe(() => {
                this.router.navigate(['/indexPayrollItem']);
            });
    }

    ngOnInit(): void {
    }
}