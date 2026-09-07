
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateBusinessUnitComponent } from './create.component';
import { BusinessUnitService } from '../../../services/BusinessUnit.service';
import { Router } from '@angular/router';

describe('CreateBusinessUnitComponent', () => {
  let component: CreateBusinessUnitComponent;
  let fixture: ComponentFixture<CreateBusinessUnitComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateBusinessUnitComponent
      ],
      providers: [
        BusinessUnitService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateBusinessUnitComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});