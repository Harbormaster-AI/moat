import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { WorkCenterService } from '../../../services/WorkCenter.service';
import { WorkCenter } from '../../../models/WorkCenter';
import { SubBaseComponent } from '../../WorkCenter/sub.base.component';

@Component({
    selector: 'app-create-workCenter',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateWorkCenterComponent extends SubBaseComponent implements OnInit {

    title = 'Add WorkCenter';

    workCenterForm: FormGroup;
    workCenter: WorkCenter;

    constructor( http: HttpClient,
        private workCenterService: WorkCenterService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.workCenterForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      capability: ['', Validators.required],
      ProductionLine: ['', ]
        });
    }

    
    addWorkCenter(name, capability, ProductionLine): void {
        this.workCenterService
        .addWorkCenter(name, capability, ProductionLine)
            .subscribe(() => {
                this.router.navigate(['/indexWorkCenter']);
            });
    }

    ngOnInit(): void {
    }
}