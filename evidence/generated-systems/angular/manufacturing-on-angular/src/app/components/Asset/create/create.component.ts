import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AssetService } from '../../../services/Asset.service';
import { Asset } from '../../../models/Asset';
import { SubBaseComponent } from '../../Asset/sub.base.component';

@Component({
    selector: 'app-create-asset',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAssetComponent extends SubBaseComponent implements OnInit {

    title = 'Add Asset';

    assetForm: FormGroup;
    asset: Asset;

    constructor( http: HttpClient,
        private assetService: AssetService,
        private fb: FormBuilder,
        private router: Router
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

    
    addAsset(assetTag, assetName, commissioningDate, Plant, WorkCenter, MaintenanceOrders, MaintenancePlans, AssetStatus): void {
        this.assetService
        .addAsset(assetTag, assetName, commissioningDate, Plant, WorkCenter, MaintenanceOrders, MaintenancePlans, AssetStatus)
            .subscribe(() => {
                this.router.navigate(['/indexAsset']);
            });
    }

    ngOnInit(): void {
    }
}