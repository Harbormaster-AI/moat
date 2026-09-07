import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AssetService } from '../../../services/Asset.service';
import { SubBaseComponent } from '../../Asset/sub.base.component';


@Component({
    selector: 'app-edit-asset',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAssetComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Asset';

    assetForm: FormGroup;
    asset: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AssetService,
        private fb: FormBuilder
) {
        super(http);
        this.assetForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  assetTag: ['', Validators.required],
      assetName: ['', Validators.required],
      commissioningDate: ['', Validators.required],
      Plant: ['', ],
      WorkCenter: ['', ],
      MaintenanceOrders: ['', ],
      MaintenancePlans: ['', ],
      AssetStatus: ['', ]
        });
    }

    
    updateAsset(assetTag, assetName, commissioningDate, Plant, WorkCenter, MaintenanceOrders, MaintenancePlans, AssetStatus): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAsset(assetTag, assetName, commissioningDate, Plant, WorkCenter, MaintenanceOrders, MaintenancePlans, AssetStatus, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAsset']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAsset(params['id']).subscribe(res => {
                this.asset = res;
            });
        });
    }
}