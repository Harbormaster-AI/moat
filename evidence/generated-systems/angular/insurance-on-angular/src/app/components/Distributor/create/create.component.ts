import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DistributorService } from '../../../services/Distributor.service';
import { Distributor } from '../../../models/Distributor';
import { SubBaseComponent } from '../../Distributor/sub.base.component';

@Component({
    selector: 'app-create-distributor',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateDistributorComponent extends SubBaseComponent implements OnInit {

    title = 'Add Distributor';

    distributorForm: FormGroup;
    distributor: Distributor;

    constructor( http: HttpClient,
        private distributorService: DistributorService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.distributorForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      licenseNumber: ['', Validators.required],
      region: ['', Validators.required],
      Insurers: ['', ],
      Agents: ['', ],
      Policies: ['', ],
      DistributorType: ['', ]
        });
    }

    
    addDistributor(name, licenseNumber, region, Insurers, Agents, Policies, DistributorType): void {
        this.distributorService
        .addDistributor(name, licenseNumber, region, Insurers, Agents, Policies, DistributorType)
            .subscribe(() => {
                this.router.navigate(['/indexDistributor']);
            });
    }

    ngOnInit(): void {
    }
}