import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PerformanceCycleService } from '../../../services/PerformanceCycle.service';
import { PerformanceCycle } from '../../../models/PerformanceCycle';
import { SubBaseComponent } from '../../PerformanceCycle/sub.base.component';

@Component({
    selector: 'app-create-performanceCycle',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePerformanceCycleComponent extends SubBaseComponent implements OnInit {

    title = 'Add PerformanceCycle';

    performanceCycleForm: FormGroup;
    performanceCycle: PerformanceCycle;

    constructor( http: HttpClient,
        private performanceCycleService: PerformanceCycleService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.performanceCycleForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      startDate: ['', Validators.required],
      endDate: ['', Validators.required],
      Organization: ['', ],
      Reviews: ['', ],
      Goals: ['', ],
      Status: ['', ]
        });
    }

    
    addPerformanceCycle(name, startDate, endDate, Organization, Reviews, Goals, Status): void {
        this.performanceCycleService
        .addPerformanceCycle(name, startDate, endDate, Organization, Reviews, Goals, Status)
            .subscribe(() => {
                this.router.navigate(['/indexPerformanceCycle']);
            });
    }

    ngOnInit(): void {
    }
}