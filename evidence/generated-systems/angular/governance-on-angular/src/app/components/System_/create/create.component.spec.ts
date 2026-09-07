
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateSystem_Component } from './create.component';
import { System_Service } from '../../../services/System_.service';
import { Router } from '@angular/router';

describe('CreateSystem_Component', () => {
  let component: CreateSystem_Component;
  let fixture: ComponentFixture<CreateSystem_Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateSystem_Component
      ],
      providers: [
        System_Service,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateSystem_Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});