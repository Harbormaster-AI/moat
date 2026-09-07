
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateEngineTypeComponent } from './create.component';
import { EngineTypeService } from '../../../services/EngineType.service';
import { Router } from '@angular/router';

describe('CreateEngineTypeComponent', () => {
  let component: CreateEngineTypeComponent;
  let fixture: ComponentFixture<CreateEngineTypeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateEngineTypeComponent
      ],
      providers: [
        EngineTypeService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateEngineTypeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});