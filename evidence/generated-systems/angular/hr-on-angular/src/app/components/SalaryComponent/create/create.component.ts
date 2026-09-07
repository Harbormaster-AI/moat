import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { SalaryComponentService } from '../../../services/SalaryComponent.service';
import { SalaryComponent } from '../../../models/SalaryComponent';
import { SubBaseComponent } from '../../SalaryComponent/sub.base.component';

@Component({
    selector: 'app-create-salaryComponent',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateSalaryComponentComponent extends SubBaseComponent implements OnInit {

    title = 'Add SalaryComponent';

    salaryComponentForm: FormGroup;
    salaryComponent: SalaryComponent;

    constructor( http: HttpClient,
        private salaryComponentService: SalaryComponentService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.salaryComponentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  amount: ['', Validators.required],
      recurring: ['', Validators.required],
      CompensationPackage: ['', ],
      ComponentType: ['', ]
        });
    }

    
    addSalaryComponent(amount, recurring, CompensationPackage, ComponentType): void {
        this.salaryComponentService
        .addSalaryComponent(amount, recurring, CompensationPackage, ComponentType)
            .subscribe(() => {
                this.router.navigate(['/indexSalaryComponent']);
            });
    }

    ngOnInit(): void {
    }
}