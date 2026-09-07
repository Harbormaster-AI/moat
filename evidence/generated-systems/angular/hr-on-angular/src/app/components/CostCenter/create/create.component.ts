import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CostCenterService } from '../../../services/CostCenter.service';
import { CostCenter } from '../../../models/CostCenter';
import { SubBaseComponent } from '../../CostCenter/sub.base.component';

@Component({
    selector: 'app-create-costCenter',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCostCenterComponent extends SubBaseComponent implements OnInit {

    title = 'Add CostCenter';

    costCenterForm: FormGroup;
    costCenter: CostCenter;

    constructor( http: HttpClient,
        private costCenterService: CostCenterService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.costCenterForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  code: ['', Validators.required],
      name: ['', Validators.required],
      Organization: ['', ],
      Departments: ['', ],
      Positions: ['', ],
      Employees: ['', ]
        });
    }

    
    addCostCenter(code, name, Organization, Departments, Positions, Employees): void {
        this.costCenterService
        .addCostCenter(code, name, Organization, Departments, Positions, Employees)
            .subscribe(() => {
                this.router.navigate(['/indexCostCenter']);
            });
    }

    ngOnInit(): void {
    }
}