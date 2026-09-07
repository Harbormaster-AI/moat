
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateInspectionCharacteristicComponent } from './create.component';
import { InspectionCharacteristicService } from '../../../services/InspectionCharacteristic.service';
import { Router } from '@angular/router';

describe('CreateInspectionCharacteristicComponent', () => {
  let component: CreateInspectionCharacteristicComponent;
  let fixture: ComponentFixture<CreateInspectionCharacteristicComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateInspectionCharacteristicComponent
      ],
      providers: [
        InspectionCharacteristicService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateInspectionCharacteristicComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});