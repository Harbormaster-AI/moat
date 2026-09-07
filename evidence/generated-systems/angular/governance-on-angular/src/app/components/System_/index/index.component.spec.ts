
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexSystem_Component } from './index.component';
import { System_Service } from '../../../services/System_.service';

describe('IndexSystem_Component', () => {
  let component: IndexSystem_Component;
  let fixture: ComponentFixture<IndexSystem_Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexSystem_Component
      ],
      providers: [
        System_Service,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexSystem_Component);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});