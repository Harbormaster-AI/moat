
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateMROFacilityComponent } from './create.component';
import { MROFacilityService } from '../../../services/MROFacility.service';
import { Router } from '@angular/router';

describe('CreateMROFacilityComponent', () => {
  let component: CreateMROFacilityComponent;
  let fixture: ComponentFixture<CreateMROFacilityComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateMROFacilityComponent
      ],
      providers: [
        MROFacilityService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateMROFacilityComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});