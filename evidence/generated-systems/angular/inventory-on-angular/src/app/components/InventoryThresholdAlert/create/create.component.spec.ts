
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateInventoryThresholdAlertComponent } from './create.component';
import { InventoryThresholdAlertService } from '../../../services/InventoryThresholdAlert.service';
import { Router } from '@angular/router';

describe('CreateInventoryThresholdAlertComponent', () => {
  let component: CreateInventoryThresholdAlertComponent;
  let fixture: ComponentFixture<CreateInventoryThresholdAlertComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateInventoryThresholdAlertComponent
      ],
      providers: [
        InventoryThresholdAlertService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateInventoryThresholdAlertComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});