
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexException_Component } from './index.component';
import { Exception_Service } from '../../../services/Exception_.service';

describe('IndexException_Component', () => {
  let component: IndexException_Component;
  let fixture: ComponentFixture<IndexException_Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexException_Component
      ],
      providers: [
        Exception_Service,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexException_Component);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});