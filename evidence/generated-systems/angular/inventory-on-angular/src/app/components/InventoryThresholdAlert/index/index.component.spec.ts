
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexInventoryThresholdAlertComponent } from './index.component';
import { InventoryThresholdAlertService } from '../../../services/InventoryThresholdAlert.service';

describe('IndexInventoryThresholdAlertComponent', () => {
  let component: IndexInventoryThresholdAlertComponent;
  let fixture: ComponentFixture<IndexInventoryThresholdAlertComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexInventoryThresholdAlertComponent
      ],
      providers: [
        InventoryThresholdAlertService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexInventoryThresholdAlertComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});