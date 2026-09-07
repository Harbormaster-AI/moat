
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAerospaceManufacturerComponent } from './create.component';
import { AerospaceManufacturerService } from '../../../services/AerospaceManufacturer.service';
import { Router } from '@angular/router';

describe('CreateAerospaceManufacturerComponent', () => {
  let component: CreateAerospaceManufacturerComponent;
  let fixture: ComponentFixture<CreateAerospaceManufacturerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAerospaceManufacturerComponent
      ],
      providers: [
        AerospaceManufacturerService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAerospaceManufacturerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});