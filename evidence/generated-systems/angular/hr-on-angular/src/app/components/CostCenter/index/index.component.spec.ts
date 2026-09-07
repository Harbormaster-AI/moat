
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCostCenterComponent } from './index.component';
import { CostCenterService } from '../../../services/CostCenter.service';

describe('IndexCostCenterComponent', () => {
  let component: IndexCostCenterComponent;
  let fixture: ComponentFixture<IndexCostCenterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCostCenterComponent
      ],
      providers: [
        CostCenterService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCostCenterComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});