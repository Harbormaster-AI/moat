
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAircraftPackageComponent } from './create.component';
import { AircraftPackageService } from '../../../services/AircraftPackage.service';
import { Router } from '@angular/router';

describe('CreateAircraftPackageComponent', () => {
  let component: CreateAircraftPackageComponent;
  let fixture: ComponentFixture<CreateAircraftPackageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAircraftPackageComponent
      ],
      providers: [
        AircraftPackageService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAircraftPackageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});