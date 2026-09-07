import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { LaboratoryService } from '../../../services/Laboratory.service';
import { Laboratory } from '../../../models/Laboratory';
import { SubBaseComponent } from '../../Laboratory/sub.base.component';

@Component({
    selector: 'app-create-laboratory',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateLaboratoryComponent extends SubBaseComponent implements OnInit {

    title = 'Add Laboratory';

    laboratoryForm: FormGroup;
    laboratory: Laboratory;

    constructor( http: HttpClient,
        private laboratoryService: LaboratoryService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.laboratoryForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      cliaNumber: ['', Validators.required],
      Facility: ['', ],
      LaboratoryOrders: ['', ],
      LabResults: ['', ]
        });
    }

    
    addLaboratory(name, cliaNumber, Facility, LaboratoryOrders, LabResults): void {
        this.laboratoryService
        .addLaboratory(name, cliaNumber, Facility, LaboratoryOrders, LabResults)
            .subscribe(() => {
                this.router.navigate(['/indexLaboratory']);
            });
    }

    ngOnInit(): void {
    }
}