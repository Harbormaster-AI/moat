
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateSerialNumberComponent } from './create.component';
import { SerialNumberService } from '../../../services/SerialNumber.service';
import { Router } from '@angular/router';

describe('CreateSerialNumberComponent', () => {
  let component: CreateSerialNumberComponent;
  let fixture: ComponentFixture<CreateSerialNumberComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateSerialNumberComponent
      ],
      providers: [
        SerialNumberService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateSerialNumberComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});