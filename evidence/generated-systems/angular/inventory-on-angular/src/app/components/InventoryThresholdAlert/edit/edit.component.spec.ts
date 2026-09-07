
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditInventoryThresholdAlertComponent } from './edit.component';
import { InventoryThresholdAlertService } from '../../../services/InventoryThresholdAlert.service';

describe('EditInventoryThresholdAlertComponent', () => {
  let component: EditInventoryThresholdAlertComponent;
  let fixture: ComponentFixture<EditInventoryThresholdAlertComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditInventoryThresholdAlertComponent
      ],
      providers: [
        InventoryThresholdAlertService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditInventoryThresholdAlertComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});