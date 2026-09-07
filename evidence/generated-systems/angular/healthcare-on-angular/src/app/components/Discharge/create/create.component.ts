import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DischargeService } from '../../../services/Discharge.service';
import { Discharge } from '../../../models/Discharge';
import { SubBaseComponent } from '../../Discharge/sub.base.component';

@Component({
    selector: 'app-create-discharge',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateDischargeComponent extends SubBaseComponent implements OnInit {

    title = 'Add Discharge';

    dischargeForm: FormGroup;
    discharge: Discharge;

    constructor( http: HttpClient,
        private dischargeService: DischargeService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.dischargeForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  dischargeDateTime: ['', Validators.required],
      Encounter: ['', ],
      Disposition: ['', ]
        });
    }

    
    addDischarge(dischargeDateTime, Encounter, Disposition): void {
        this.dischargeService
        .addDischarge(dischargeDateTime, Encounter, Disposition)
            .subscribe(() => {
                this.router.navigate(['/indexDischarge']);
            });
    }

    ngOnInit(): void {
    }
}