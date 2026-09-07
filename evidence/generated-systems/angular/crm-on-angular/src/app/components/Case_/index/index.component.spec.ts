
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCase_Component } from './index.component';
import { Case_Service } from '../../../services/Case_.service';

describe('IndexCase_Component', () => {
  let component: IndexCase_Component;
  let fixture: ComponentFixture<IndexCase_Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCase_Component
      ],
      providers: [
        Case_Service,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCase_Component);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});